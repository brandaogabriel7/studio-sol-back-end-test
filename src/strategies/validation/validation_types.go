package validation

type ValidationType string

const MIN_SIZE ValidationType = "minSize"
const MIN_DIGIT ValidationType = "minDigit"
const MIN_SPECIAL_CHARS ValidationType = "minSpecialChars"
const NO_REPETED ValidationType = "noRepeted"