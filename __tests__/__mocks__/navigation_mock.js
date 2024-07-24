/*
 * Mock navigation data for frontend tests
 * To edit the actual menu, see data/navigation.yaml
 */

/* eslint-disable import/no-default-export */
export default [
  {
    title: "Use GitLab",
    url: "user.html",
    submenu: [
      {
        title: "Learn Git",
        url: "topics/git.html",
        submenu: [
          {
            title: "Getting started",
            url: "topics/git/get_started.html",
          },
          {
            title: "Install and configure Git",
            url: "topics/git/how_to_install_git.html",
          },
          {
            title: "Tutorial: Make your first Git commit",
            url: "tutorials/make_first_git_commit.html",
          },
          {
            title: "Clone a Git repository",
            url: "topics/git/clone.html",
            submenu: [
              {
                title: "Use partial clones to reduce clone size",
                url: "topics/git/partial_clone.html",
              },
            ],
          },
          {
            title: "Create a Git branch",
            url: "topics/git/branch.html",
          },
          {
            title: "Add files and make changes",
            url: "gitlab-basics/add-file.html",
          },
          {
            title: "Tutorial: Update Git commit messages",
            url: "tutorials/update_commit_messages.html",
          },
          {
            title: "Use additional Git commands",
            url: "gitlab-basics/start-using-git.html",
          },
          {
            title: "Rebase and force-push",
            url: "topics/git/git_rebase.html",
          },
          {
            title: "Feature branch workflow",
            url: "gitlab-basics/feature_branch_workflow.html",
          },
          {
            title: "Undo changes",
            url: "topics/git/undo.html",
          },
          {
            title: "Push options",
            url: "user/project/push_options.html",
          },
          {
            title: "Troubleshooting",
            url: "topics/git/troubleshooting_git.html",
          },
        ],
      },
    ],
  },
  {
    title: "About",
    url: "devsecops/",
    submenu: [
      {
        title: "Use of generative AI",
        url: "legal/use_generative_ai.html",
      },
    ],
  },
  {
    title: "GitLab Design System",
    url: "https://design.gitlab.com",
  },
];
