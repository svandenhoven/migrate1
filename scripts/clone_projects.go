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
 *    branch, see the ProductCloneInfo method).
 */
func (p Product) ShouldSkipClone(branch string) bool {
	return branch == p.DefaultBranch &&
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
func (p Product) CloneInfo(productName string) (string, string) {
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
		branchName = p.DefaultBranch
	}
	if mergeRequestIID == "" {
		return branchName, fmt.Sprintf("heads/%s", branchName)
	}

	return branchName, fmt.Sprintf("merge-requests/%s/head", mergeRequestIID)
}

func (p Product) Clone(productName string) error {
	cloneDir := p.CloneDir
	branch, refspec := p.CloneInfo(productName)

	// Limit the pipeline to pull only the repo where the MR is to save time and space.
	if p.ShouldSkipClone(branch) {
		log.Printf("[Info] Skipping %s", productName)

		return nil
	}

	// Remove the product repo if it already exists if REMOVE_BEFORE_CLONE is true,
	// or if we're in a CI environment.
	// This can happen if we land on a Runner that already ran a docs build.
	if _, err := os.Stat(cloneDir); err == nil && (os.Getenv("CI") == "true" || os.Getenv("REMOVE_BEFORE_CLONE") == "true") {
		err = os.RemoveAll(cloneDir)
		if err != nil {
			return fmt.Errorf("Error removing directory: %v\n", err)
		}

		log.Printf("[Info] %s already exists, removing it", cloneDir)
	}

	// If the directory exists, and it's a local environment, skip it.
	if _, err := os.Stat(cloneDir); err == nil && os.Getenv("CI") == "" {
		log.Printf("[Info] %s directory already exists, skipping", productName)

		return nil
	}

	// Create the target directory, and move into it
	if err := os.MkdirAll(cloneDir, os.ModePerm); err != nil {
		return fmt.Errorf("Error creating directory: %v\n", err)
	} else if err := os.Chdir(cloneDir); err != nil {
		return fmt.Errorf("Error changing directory: %v\n", err)
	}

	// Initialize the repository, and fetch the desired branch and refspec
	log.Printf("[Info] Fetching %s on branch %s at commit %s", productName, branch, refspec)
	runGitCommand("-c", fmt.Sprintf("init.defaultBranch=%s", branch), "init")
	runGitCommand("remote", "add", "origin", p.Repo)
	runGitCommand("fetch", "--depth", "1", "origin", refspec)
	runGitCommand("-c", "advice.detachedHead=false", "checkout", "FETCH_HEAD")

	// Print the last commit message
	logMessage, err := runGitCommand("log", "--oneline", "-n", "1")
	if err != nil {
		return fmt.Errorf("Error reading Git log: %v\n", err)
	}

	log.Printf("[Info] Last commit: %s", string(logMessage))

	return nil
}

func main() {
	// Load product info
	productsData, err := readProductData("data/products.yaml")
	if err != nil {
		log.Fatalf("Error loading products.yaml: %v\n", err)
	}

	// Iterate through products
	for productName, product := range productsData.Products {
		err := product.Clone(productName)
		if err != nil {
			log.Fatal(err)
		}
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

func runGitCommand(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error running Git command 'git %v': %v\n", args, err)

		return "", err
	}

	return string(output), nil
}
