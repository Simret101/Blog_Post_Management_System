
# Git Hooks for Blog Post Management System

This repository contains a collection of Git hooks that enforce a clean commit policy for your **Blog Post Management System**.

### The Hooks:
1. **pre-commit**: Prevents committing files on the `master` branch.
2. **commit-msg**: Ensures commit messages follow the defined rules for consistency.
3. **.git-commit-msg-template.txt**: A default template for your commit messages.

---

## Installation

### 1. Clone This Repository
Clone the repository to your local device to get started.

```bash
git clone https://github.com/Simret101/Blog_Post_Management_System.git
cd Blog_Post_Management_System
```

---

### 2. Commit Message Template
To use the commit message template, follow these steps:

- Copy the `.git-commit-msg-template.txt` file to a location of your choice (e.g., your home directory).

  ```bash
  cp .git-commit-msg-template.txt ~/.git-commit-msg-template.txt
  ```

- Set the template as the default for your commit messages by running:

  ```bash
  git config --global commit.template ~/.git-commit-msg-template.txt
  ```

Now, every time you run `git commit`, the template will appear, guiding you to follow consistent commit message conventions.

*Tip:* If you're using PhpStorm as your Git client, you can use the plugin to automatically use commit templates.

---

### 3. Installing Hooks

You can install the hooks either for a **single repository** or globally across all repositories.

#### **Single Repository Installation**
If you want to use the hooks only in your Blog Post Management System repository:

1. Copy the `pre-commit` and `commit-msg` hook files to the `.git/hooks` directory of your repository:

   ```bash
   cp {pre-commit,commit-msg} .git/hooks/
   chmod +x .git/hooks/{pre-commit,commit-msg}
   ```

#### **Global Installation**
To install the hooks globally, which will apply to all repositories:

1. Create a folder to hold the hooks:

   ```bash
   mkdir -p ~/.git-template/hooks
   ```

2. Configure Git to use the new template directory globally:

   ```bash
   git config --global init.templatedir '~/.git-template'
   ```

3. Copy the `pre-commit` and `commit-msg` hooks into the global hooks folder:

   ```bash
   cp {pre-commit,commit-msg} ~/.git-template/hooks
   chmod +x ~/.git-template/hooks/{pre-commit,commit-msg}
   ```

4. From now on, every time you run `git init` or `git clone`, the hooks will automatically be applied. To apply the hooks to an existing repository, simply run:

   ```bash
   git init
   ```

---

## Configuration

You can further customize your Git hooks behavior by modifying the Git configuration.

### 1. Append Branch Name to Commit Messages
To automatically append the current branch name to commit messages, run:

```bash
git config git-hooks.append-branch-name true
```

By default, this option is disabled. If you don't want the branch name appended, set it to `false`.

### 2. Enable Committing on the `master` Branch
By default, committing on the `master` branch is disabled. To allow commits on `master`, run:

```bash
git config git-hooks.commit-on-master true
```

If you'd like to disable committing on the `master` branch again, run:

```bash
git config git-hooks.commit-on-master false
```

---

### Global Configuration
To set up these configurations globally, use the `--global` flag:

```bash
git config --global git-hooks.append-branch-name true
git config --global git-hooks.commit-on-master true
```

---
