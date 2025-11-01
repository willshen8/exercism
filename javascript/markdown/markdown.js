const wrap = (text, tag) => {
  return `<${tag}>${text}</${tag}>`;
};

// removed unused function isTag

const parser = (markdown, delimiter, tag) => {
  const pattern = new RegExp(`${delimiter}(.+)${delimiter}`);
  const replacement = `<${tag}>$1</${tag}>`;
  return markdown.replace(pattern, replacement);
};

const parseDoubleUnderscore = (markdown) => {
  return parser(markdown, '__', 'strong');
};

const parseSingleUnderscore = (markdown) => {
  return parser(markdown, '_', 'em');
};

const parseText = (markdown, list) => {
  const parsedText = parseSingleUnderscore(parseDoubleUnderscore(markdown));
  if (!list) {
    return wrap(parsedText, 'p');
  }
  return parsedText;
};

const parseHeader = (markdown, list) => {
  let headerLevels = 0;

  for (let i = 0; i < markdown.length; i++) {
    if (markdown[i] === '#') {
      headerLevels += 1;
    } else {
      break;
    }
  }
  if (headerLevels === 0 || headerLevels > 6) {
    return [null, list];
  }
  const headerTag = `h${headerLevels}`;
  const headerHtml = wrap(markdown.substring(headerLevels + 1), headerTag);
  if (!list) {
    return [headerHtml, false];
  }
  return [`</ul>${headerHtml}`, false];
};

const parseLineItem = (markdown, list) => {
  if (!markdown.startsWith('*')) return [null, list];

  const innerHtml = wrap(parseText(markdown.substring(2), true), 'li');
  if (list && markdown.startsWith('*')) {
    return [innerHtml, true];
  }
  return [`<ul>${innerHtml}`, true];
};

const parseParagraph = (markdown, list) => {
  if (!list) {
    return [parseText(markdown, false), false];
  }
  return [`</ul>${parseText(markdown, false)}`, false];
};

const parseLine = (markdown, list) => {
  let [result, inListAfter] = parseHeader(markdown, list);
  if (result === null) {
    [result, inListAfter] = parseLineItem(markdown, list);
  }
  if (result === null) {
    [result, inListAfter] = parseParagraph(markdown, list);
  }
  if (result === null) {
    throw new Error('Remove this line and implement the function');
  }
  return [result, inListAfter];
};

export const parse = (markdown) => {
  const lines = markdown.split('\n');
  let result = '';
  let isList = false;
  for (let i = 0; i < lines.length; i++) {
    let [lineResult, newList] = parseLine(lines[i], isList);
    result += lineResult;
    isList = newList;
  }
  if (isList) {
    return result + '</ul>';
  }
  return result;
};
