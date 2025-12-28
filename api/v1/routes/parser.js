const fs = require('fs');
const path = require('path');
const { parse } = require('csv-parse/sync');

class Parser {
  constructor(config = {}) {
    this.config = {
      delimiter: ',',
      skipEmptyLines: true,
      columns: true,
      ...config,
    };
  }

  parseFile(filePath) {
    try {
      const absolutePath = path.resolve(filePath);
      const fileContent = fs.readFileSync(absolutePath, 'utf-8');
      return parse(fileContent, this.config);
    } catch (error) {
      throw new Error(`Failed to parse file: ${error.message}`);
    }
  }

  parseString(content) {
    try {
      return parse(content, this.config);
    } catch (error) {
      throw new Error(`Failed to parse string: ${error.message}`);
    }
  }
}

module.exports = Parser;