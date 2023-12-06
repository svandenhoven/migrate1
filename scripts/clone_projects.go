package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v3"
)

// A Product is a GitLab project that we include on the Docs site
type Product struct {
	Repo          string `yaml:"repo"`
	DefaultBranch string `yaml:"default_branch"`
	CloneDir      string `yaml:"clone_dir"`
}

type Products struct {
	Products map[string]Product `yaml:"products"`
}

func main() {
	// Load product info
	productsData, err := readProductData("data/products.yaml")
	if err != nil {
		log.Fatalf("Error loading products.yaml: %v\n", err)
	}

	// Iterate through products
	for name, product := range productsData.Products {
		cloneDir := product.CloneDir
		branch, refspec := productCloneInfo(name, product)

		// Limit the pipeline to pull only the repo where the MR is to save time and space.
		if shouldSkipClone(branch, product.DefaultBranch) {
			log.Printf("[Info] Skipping %s", name)
			continue
		}

		// Remove the product repo if it already exists if REMOVE_BEFORE_CLONE is true,
		// or if we're in a CI environment.
		// This can happen if we land on a Runner that already ran a docs build.
		if _, err := os.Stat(cloneDir); err == nil && (os.Getenv("CI") == "true" || os.Getenv("REMOVE_BEFORE_CLONE") == "true") {
			err = os.RemoveAll(cloneDir)
			if err != nil {
				log.Fatalf("Error removing directory: %v\n", err)
			}
			log.Printf("[Info] %s already exists, removing it", cloneDir)
		}
		// If the directory exists, and it's a local environment, skip it.
		if _, err := os.Stat(cloneDir); err == nil && os.Getenv("CI") == "" {
			log.Printf("[Info] %s directory already exists, skipping", name)
			continue
		}

		// Create the target directory, and move into it
		if err := os.MkdirAll(cloneDir, os.ModePerm); err != nil {
			log.Fatalf("Error creating directory: %v\n", err)
		} else if err := os.Chdir(cloneDir); err != nil {
			log.Fatalf("Error changing directory: %v\n", err)
		}

		// Initialize the repository, and fetch the desired branch and refspec
		log.Printf("[Info] Fetching %s on branch %s at commit %s", name, branch, refspec)
		runGitCommand("git", "-c", fmt.Sprintf("init.defaultBranch=%s", branch), "init")
		runGitCommand("git", "remote", "add", "origin", product.Repo)
		runGitCommand("git", "fetch", "--depth", "1", "origin", refspec)
		runGitCommand("git", "-c", "advice.detachedHead=false", "checkout", "FETCH_HEAD")

		// Print the last commit message
		logCmd := exec.Command("git", "log", "--oneline", "-n", "1")
		logMessage, err := logCmd.CombinedOutput()
		if err != nil {
			log.Fatalf("Error reading Git log: %v\n", err)
		}
		log.Printf("[Info] Last commit: %s", string(logMessage))
	}
}

/**
 * Return product info from a given YAML file
 */
func readProductData(productsYaml string) (*Products, error) {
	data, err := os.ReadFile(productsYaml)
	if err != nil {
		return nil, err
	}
	var productsData Products
	err = yaml.Unmarshal(data, &productsData)
	return &productsData, err
}

/**
 * Determine if a project clone can be skipped.
 *
 * The following is used only in review apps triggered from one of the five
 * products. It limits the pipeline to pull only the repo where the MR is, not
 * all five, to save time. If ALL of the following are true, skip the
 * clone (remember, this runs in gitlab-docs):
 *
 * 1. If the pipeline was triggered via the API (multi-project pipeline)
 *    (to exclude the case where we create a branch off gitlab-docs)
 * 2. If the remote branch is the upstream's product default branch name
 *    (which means BRANCH_<slug> is missing, so we default to the default
 *    branch, see the productCloneInfo method).
 */
func shouldSkipClone(branch string, defaultBranch string) bool {
	return branch == defaultBranch &&
		os.Getenv("CI_PIPELINE_SOURCE") == "trigger"
}

/**
 * Determine the Git branch and refspec to fetch.
 *
 * Use the BRANCH_* environment variable, and if not assigned,
 * set to the default branch.
 *
 * @todo Use the project's stable branch rather than the default.
 * https://gitlab.com/gitlab-org/technical-writing-group/gitlab-docs-hugo/-/issues/51
 */
func productCloneInfo(productName string, product Product) (string, string) {
	/**
	 * The BRANCH_ and MERGE_REQUEST_IID_ variables that we use here come from
	 * https://gitlab.com/gitlab-org/gitlab/-/blob/master/scripts/trigger-build.rb.
	 * Usually they match the product name, but the GitLab product uses "EE".
	 */
	envVarSuffix := func() string {
		if productName == "gitlab" {
			return "EE"
		}
		return strings.ToUpper(productName)
	}()

	mergeRequestIID := os.Getenv(fmt.Sprintf("MERGE_REQUEST_IID_%s", envVarSuffix))
	branchName := os.Getenv(fmt.Sprintf("BRANCH_%s", envVarSuffix))

	if branchName == "" {
		branchName = product.DefaultBranch
	}
	if mergeRequestIID == "" {
		return branchName, fmt.Sprintf("heads/%s", branchName)
	}

	return branchName, fmt.Sprintf("merge-requests/%s/head", mergeRequestIID)
}

func runGitCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error running Git command '%s': %v\n", command, err)
	}
}
