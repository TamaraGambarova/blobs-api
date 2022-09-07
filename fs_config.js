const operations = require("./operations");

module.exports = {
  dirsToCreate: {
    common: {
      render: operations.defaultDirCreator([
        "internal",
        "internal/config",
        "internal/cli"
      ])
    },
    db: {
      render: operations.defaultDirCreator([
        "internal/assets",
        "internal/assets/migrations"
      ])
    },
    http: {
      render: operations.defaultDirCreator([
        "internal/service/handlers",
        "internal/service/requests"
      ])
    },
    docs: {
      render: operations.defaultDirCreator([
        "docs",
        "docs/spec",
        "docs/web",
        "docs/spec/components",
        "docs/spec/components/parameters",
        "docs/spec/components/schemas",
        "docs/spec/paths"
      ])
    },
    render: function(generator) {
      this.common.render(generator);
      this.db.render(generator);
      this.http.render(generator);
      this.docs.render(generator);
    }
  },
  filesToCopy: {
    common: {
      render: operations.defaultFileCopier([
        {
          input: ".gitignore",
          output: "..gitignore"
        }
      ])
    },
    db: {
      render: operations.defaultFileCopier([
        {
          input: "internal/assets/main.go",
          output: "internal/assets/main.go"
        },
        {
          input: "internal/assets/migrations/001_initial.sql",
          output: "internal/assets/migrations/001_initial.sql"
        }
      ])
    },
    http: {
      render: operations.defaultFileCopier([
        {
          input: "internal/service/handlers/ctx.go",
          output: "internal/service/handlers/ctx.go"
        }
      ])
    },
    docs: {
      render: operations.defaultFileCopier([
        {
          input: "docs/spec/components/parameters/pageLimitParam.yaml",
          output: "docs/spec/components/parameters/pageLimitParam.yaml"
        },
        {
          input: "docs/spec/components/parameters/pageNumberParam.yaml",
          output: "docs/spec/components/parameters/pageNumberParam.yaml"
        },
        {
          input: "docs/spec/components/parameters/sortingParam.yaml",
          output: "docs/spec/components/parameters/sortingParam.yaml"
        },
        {
          input: "docs/spec/components/README.md",
          output: "docs/spec/components/README.md"
        },
        {
          input: "docs/spec/README.md",
          output: "docs/spec/README.md"
        },
        {
          input: "docs/web/favicon.png",
          output: "docs/web/favicon.png"
        },
        {
          input: "docs/web/redoc-config.yaml",
          output: "docs/web/redoc-config.yaml"
        },
        {
          input: "docs/LICENSE",
          output: "docs/LICENSE"
        }
      ])
    },
    render: function(generator) {
      this.common.render(generator);
      this.db.render(generator);
      this.http.render(generator);
      this.docs.render(generator);
    }
  },
  filesToRender: {
    common: {
      render: operations.defaultFileRenderer([
        {
          input: "README.md",
          output: "README.md"
        },
        {
          input: "Dockerfile",
          output: "Dockerfile"
        },
        {
          input: "main.go",
          output: "main.go"
        },
        {
          input: "internal/cli/main.go",
          output: "internal/cli/main.go"
        },
        {
          input: "internal/config/main.go",
          output: "internal/config/main.go"
        },
        {
          input: "config.yaml",
          output: "config.yaml"
        },
        {
          input: "internal/service/main.go",
          output: "internal/service/main.go"
        }
      ])
    },
    db: {
      render: operations.defaultFileRenderer([
        {
          input: "internal/cli/migrate.go",
          output: "internal/cli/migrate.go"
        }
      ])
    },
    http: {
      render: operations.defaultFileRenderer([
        {
          input: "internal/service/router.go",
          output: "internal/service/router.go"
        }
      ])
    },
    docs: {
      render: operations.defaultFileRenderer([
        {
          input: "docs/spec/openapi.yaml",
          output: "docs/spec/openapi.yaml"
        },
        {
          input: "docs/redoclyrc",
          output: "docs/.redoclyrc"
        },
        {
          input: "docs/web/index.html",
          output: "docs/web/index.html"
        },
        {
          input: "docs/web/index.hbs",
          output: "docs/web/index.hbs"
        },
        {
          input: "generate.sh",
          output: "generate.sh"
        },
        {
          input: "docs/package.json",
          output: "docs/package.json"
        },
        {
          input: "docs/package-lock.json",
          output: "docs/package-lock.json"
        }
      ])
    },
    render: function(generator) {
      this.common.render(generator);
      this.db.render(generator);
      this.http.render(generator);
      this.docs.render(generator);
    }
  },
  gitlabCI: {
    common: {
      render: operations.defaultFileRenderer([
        {
          input: "gitlab-ci.yml",
          output: ".gitlab-ci.yml"
        }
      ])
    },
    render: function(generator) {
      this.common.render(generator);
    }
  },
  travisCI: {
    common: {
      render: operations.defaultFileRenderer([
        {
          input: "travis.yml",
          output: ".travis.yml"
        }
      ])
    },
    render: function(generator) {
      this.common.render(generator);
    }
  },
  render: function(generator) {
    this.dirsToCreate.render(generator);
    this.filesToCopy.render(generator);
    this.filesToRender.render(generator);
    this.travisCI.render(generator);
    this.gitlabCI.render(generator);
  }
};
