import parse from './extract.js'
import {jest} from '@jest/globals';


test('simple_query', () => {
  expect(parse(`
select * from foo
  -- Sigma Σ {"kind":"adhoc","request-id":"e0bd080279a9e63","email":"robert@sundeck.io"}
`).length).toBe(1);
});
