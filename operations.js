const mkdirp = require("mkdirp");

module.exports = {
  nop: nop,
  defaultFileRenderer: defaultFileRenderer,
  defaultFileCopier: defaultFileCopier,
  defaultDirCreator: defaultDirCreator
};

function nop() {}

function defaultFileRenderer(files) {
  return generator => {
    files.forEach(file => {
      generator.log(`Rendering file ${file.input}`);
      generator.fs.copyTpl(
        generator.templatePath(file.input),
        generator.destinationPath(file.output),
        generator.props
      );
    });
  };
}

function defaultFileCopier(files) {
  return generator => {
    files.forEach(file => {
      generator.log(`Copying file ${file.input}`);
      generator.fs.copy(
        generator.templatePath(file.input),
        generator.destinationPath(file.output)
      );
    });
  };
}

function defaultDirCreator(directories) {
  return generator => {
    directories.forEach(item => {
      generator.log(`Creating directory ${item}`);
      mkdirp(item);
    });
  };
}
