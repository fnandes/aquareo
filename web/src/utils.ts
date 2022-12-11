
const metricNames = new Map<string, string>([
  ['ca', 'Calcium'],
  ['mg', 'Magnesium'],
  ['no2', 'Nitrite'],
  ['no3', 'Nitrate'],
  ['po4', 'Phosphate'],
  ['co2', 'Carbon Dioxide'],
  ['nh4', 'Anmonia']
])

export const getMetricName = (formula: string): string => {
  const name = metricNames.get(formula)
  return name ? `${name} (${formula.toUpperCase()})` : formula.toUpperCase()
}