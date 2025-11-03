const wrap = (text, tag) => `<${tag}>${text}</${tag}>`;

const parser = (markdown, delimiter, tag) => {
  const pattern = new RegExp(`${delimiter}(.+)${delimiter}`);
  const replacement = `<${tag}>$1</${tag}>`;
  return markdown.replace(pattern, replacement);
};

const delimitersWithTags = Object.freeze([
  { delimiter: '__', tag: 'strong' },
  { delimiter: '_', tag: 'em' },
]);

const parseText = (markdown) => {
  delimitersWithTags.forEach((_, index) => {
    const { delimiter, tag } = delimitersWithTags[index];
    markdown = parser(markdown, delimiter, tag);
  });

  return markdown;
};

const parseParagraph = (markdown) => wrap(parseText(markdown), 'p');

const parseHeader = (markdown) => {
  let headerLevels = 0;

  for (let i = 0; i < markdown.length; i++) {
    if (markdown[i] === '#') {
      headerLevels += 1;
    } else {
      break;
    }
  }
  if (headerLevels === 0 || headerLevels > 6) {
    return null;
  }
  const headerTag = `h${headerLevels}`;
  return wrap(markdown.substring(headerLevels + 1), headerTag);
};

const isListItem = (markdown) => markdown.startsWith('*');

const parseLineItem = (markdown) => {
  if (!isListItem(markdown)) return null;
  const innerHtml = wrap(parseText(markdown.substring(2), true), 'li');
  return markdown.startsWith('*') ? innerHtml : `<ul>${innerHtml}`;
};

const parseLine = (markdown) => {
  let result = parseHeader(markdown);
  if (result === null) {
    result = parseLineItem(markdown);
  }
  if (result === null) {
    result = parseParagraph(markdown);
  }
  return result;
};

export const parse = (markdown) => {
  const lines = markdown.split('\n');
  let result = '';
  let isList = false;
  let isFirstListItem = true;
  for (let i = 0; i < lines.length; i++) {
    isList = isListItem(lines[i]);
    let lineResult = parseLine(lines[i]);
    if (!isFirstListItem && !isList) {
      lineResult = `</ul>${lineResult}`;
    }
    if (isList && isFirstListItem) {
      lineResult = `<ul>${lineResult}`;
      isFirstListItem = false;
    }
    result += lineResult;
  }

  return isList ? result + '</ul>' : result;
};
