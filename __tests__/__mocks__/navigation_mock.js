/*
 * Mock navigation data for frontend tests
 * To edit the actual menu, see data/navigation.yaml
 */

/* eslint-disable import/no-default-export */
export default [
  {
    title: "Use GitLab",
    url: "user/",
    submenu: [
      {
        title: "Learn Git",
        url: "topics/git/",
        submenu: [
          {
            title: "Getting started",
            url: "topics/git/get_started/",
          },
          {
            title: "Install and configure Git",
            url: "topics/git/how_to_install_git/",
          },
          {
            title: "Tutorial: Make your first Git commit",
            url: "tutorials/make_first_git_commit/",
          },
          {
            title: "Clone a Git repository",
            url: "topics/git/clone/",
            submenu: [
              {
                title: "Use partial clones to reduce clone size",
                url: "topics/git/partial_clone/",
              },
            ],
          },
          {
            title: "Create a Git branch",
            url: "topics/git/branch/",
          },
          {
            title: "Add files and make changes",
            url: "gitlab-basics/add-file/",
          },
          {
            title: "Tutorial: Update Git commit messages",
            url: "tutorials/update_commit_messages/",
          },
          {
            title: "Use additional Git commands",
            url: "gitlab-basics/start-using-git/",
          },
          {
            title: "Rebase and force-push",
            url: "topics/git/git_rebase/",
          },
          {
            title: "Feature branch workflow",
            url: "gitlab-basics/feature_branch_workflow/",
          },
          {
            title: "Undo changes",
            url: "topics/git/undo/",
          },
          {
            title: "Push options",
            url: "user/project/push_options/",
          },
          {
            title: "Troubleshooting",
            url: "topics/git/troubleshooting_git/",
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
        url: "legal/use_generative_ai/",
      },
    ],
  },
  {
    title: "GitLab Design System",
    url: "https://design.gitlab.com",
  },
];
