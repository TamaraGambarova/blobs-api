"use strict";

const Generator = require("yeoman-generator");
const chalk = require("chalk");
const yosay = require("yosay");
const fsConfig = require("./fs_config");
const operations = require("./operations");

module.exports = class extends Generator {
  prompting() {
    // Have Yeoman greet the user.
    this.log(
      yosay(
        `Welcome to the super ${chalk.red(
          "generator-tokend-module"
        )} generator!`
      )
    );

    var prompts = [
      {
        name: "serviceName",
        message: "What is the name of new service?"
      },
      {
        type: "list",
        name: "gitService",
        message: "Which git hosting are you going to use?",
        choices: ["GitLab", "GitHub", "Other"]
      },
      {
        type: "list",
        name: "imagePublished",
        message: "Where do you want to publish the image?",
        choices: ["Dockerhub", "Private Gitlab registry"]
      },
      {
        type: "confirm",
        name: "useDB",
        message: "Would you like to use PostgreSQL database for this service?",
        default: false
      },
      {
        type: "confirm",
        name: "handleHTTP",
        message: "Would you like this service to handle HTTP requests?",
        default: false
      },
      {
        type: "confirm",
        name: "hasDocs",
        message: "Would you have documentation pages for this service?",
        default: false
      },
      {
        name: "packageName",
        message: "What is the package name for the new service?"
      }
    ];

    return this.prompt(prompts).then(props => {
      this.props = props;

      if (!this.props.hasDocs) {
        fsConfig.dirsToCreate.docs.render = operations.nop;
        fsConfig.filesToCopy.docs.render = operations.nop;
        fsConfig.filesToRender.docs.render = operations.nop;
      }

      if (!this.props.useDB) {
        fsConfig.dirsToCreate.db.render = operations.nop;
        fsConfig.filesToCopy.db.render = operations.nop;
        fsConfig.filesToRender.db.render = operations.nop;
      }

      if (!this.props.handleHTTP) {
        fsConfig.dirsToCreate.http.render = operations.nop;
        fsConfig.filesToCopy.http.render = operations.nop;
        fsConfig.filesToRender.http.render = operations.nop;
      }

      switch (this.props.gitService) {
        case "GitLab":
          fsConfig.travisCI.render = operations.nop;
          break;
        case "GitHub":
          fsConfig.gitlabCI.render = operations.nop;
          break;
        default:
          fsConfig.gitlabCI.render = operations.nop;
          fsConfig.travisCI.render = operations.nop;
      }

      if (this.props.packageName === "") {
        this.props.packageName = "gitlab.com/tokend/" + this.props.serviceName;
      }
    });
  }

  writing() {
    fsConfig.render(this);
  }

  install() {
    this.log("Installing goimports if not already installed");
    this.spawnCommandSync("go", [
      "install",
      "golang.org/x/tools/cmd/goimports@latest"
    ]);
    this.log(
      "Formatting code and updating import lines with goimports in files:"
    );
    this.spawnCommandSync("goimports", ["-w", "-l", "."]);
  }

  end() {
    this.log("");
    this.log(chalk.green("Generated!"));
    if (this.props.gitService === "GitHub") {
      this.log(chalk.yellow("NOTICE:"));
      this.log(
        "I've generated .travis.yml for you, but it is incomplete there are several steps to enable Travis CI:\n\t1. Create repository for this project on GitHub;\n\t2. Log in to https://travis-ci.org and switch on CI for this repository;\n\t3. Run following commands in the root of the repository:\n\t\ttravis encrypt DOCKERHUB_USERNAME=<username> --add env.global\n\t\ttravis encrypt DOCKERHUB_PASSWORD=<password> --add env.global\n\t\ttravis encrypt GITLABREG_USERNAME=<username> --add env.global\n\t\ttravis encrypt GITLABREG_PASSWORD=<password> --add env.global\n\t4. Commit your new .travis.yml\n\t5. ...\n\t6. You are breathtaking!"
      );
      if (this.props.hasDocs) {
        this.log(chalk.yellow("NOTICE:"));
        this.log(
          "You've said that you are using GitHub and want to have GitHub Pages site. I've generated docs/ folder, but you'll need manually set up your GitHub repository for using GitHub Pages and you'll also have to build index.html manually each time."
        );
      }
    }

    if (this.props.gitService === "GitLab") {
      this.log("");
      this.log(chalk.yellow("NOTICE:"));
      this.log(
        "I've generated .gitlab-ci.yml for you, but you'll have to set env variables $DOCKERHUB_USER and $DOCKERHUB_PWD in your GitLab repository and only then foxy will be able to push release images to your docker hub. "
      );
      this.log("");
    }

    this.log("");
    this.log(chalk.green("Thank you for coming. I was glad to help you =)"));
  }
};
