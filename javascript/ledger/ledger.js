export class LedgerEntry {
  constructor(dateString, description, change) {
    this.date = new Date(dateString);
    this.description = description;
    this.change = change;
  }
}

export const createEntry = (...args) => new LedgerEntry(...args);

const formatOptionFromLocale = (currency) => ({
  currency: currency,
  style: 'currency',
  currencyDisplay: 'narrowSymbol',
  minimumFractionDigits: 2,
  maximumFractionDigits: 2,
});

const TABLE_HEADER_LOCALE = [
  {
    Locale: 'en-US',
    Date: 'Date',
    Description: 'Description',
    Change: 'Change',
  },
  {
    Locale: 'nl-NL',
    Date: 'Datum',
    Description: 'Omschrijving',
    Change: 'Verandering',
  },
];

const generateTableHeader = (locale) => {
  const { Date, Description, Change } = TABLE_HEADER_LOCALE.find(
    (localeItem) => locale === localeItem.Locale
  );

  return (
    `${Date}`.padEnd(10, ' ') +
    ' | ' +
    `${Description}`.padEnd(25, ' ') +
    ' | ' +
    `${Change}`.padEnd(13, ' ') +
    '\n'
  );
};

const formatDateString = (locale, entry) => {
  switch (locale) {
    case 'en-US':
      return `${(entry.date.getMonth() + 1)
        .toString()
        .padStart(2, '0')}/${entry.date
        .getDate()
        .toString()
        .padStart(2, '0')}/${entry.date.getFullYear()}`;
    case 'nl-NL':
      return `${entry.date.getDate().toString().padStart(2, '0')}-${(
        entry.date.getMonth() + 1
      )
        .toString()
        .padStart(2, '0')}-${entry.date.getFullYear()}`;
    default:
      throw new Error('Invalid locale, unable to format DateString');
  }
};

const truncateDescription = (entry) => {
  return entry.description.length > 25
    ? `${entry.description.substring(0, 22)}...`
    : entry.description.padEnd(25, ' ');
};

const formatChangeString = (currency, locale, entry) => {
  const formattingOptions = formatOptionFromLocale(currency);
  const changeStr = `${(entry.change / 100).toLocaleString(
    locale,
    formattingOptions
  )} `;
  switch (locale) {
    case 'en-US':
      return entry.change < 0
        ? `(${Math.abs(entry.change / 100).toLocaleString(
            locale,
            formattingOptions
          )})`
        : changeStr;
    case 'nl-NL':
      return changeStr;
    default:
      throw new Error('Invalid locale, unable to process format Change');
  }
};

export const formatEntries = (currency, locale, entries) => {
  let table = generateTableHeader(locale);
  // Sort entries
  entries.sort(
    (a, b) =>
      a.date - b.date ||
      a.change - b.change ||
      a.description.localeCompare(b.description)
  );

  entries.forEach((entry) => {
    // Write entry date to table
    table += `${formatDateString(locale, entry)} | `;

    // Write entry description to table
    table += `${truncateDescription(entry)} | `;

    // Write entry change to table
    table +=
      formatChangeString(currency, locale, entry).padStart(13, ' ') + '\n';
  });

  return table.replace(/\n$/, '');
};
