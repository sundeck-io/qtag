import extract from './extract.js'
import antlr from './antlrparser.js'
import hand from './handparser.js'
import fs from 'fs'
import yaml from 'js-yaml'


const tests = yaml.load(fs.readFileSync('./src/qtag/tests/test.yaml', 'utf8'));

for (const p of [{parser: antlr, name: "antlr"}, {parser: hand, name: "hand"}]) {
  for (const t of tests) {
    test(p.name + " - " + t.name, () => {
      const vals = extract(p.parser, t.query, t.all, t.allattributes);
      let outcome = [];
      for (const tag of t.outcome) {
        let newTag = {}
        for (const key in tag) {
          newTag[key.toUpperCase()] = tag[key];
        }
        outcome.push(newTag);
      }
      expect(vals).toStrictEqual(outcome);
    });
  }
}

