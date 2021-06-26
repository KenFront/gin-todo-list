export const validateAccount = (value: string) => {
  switch (true) {
    case !value:
      return `Accounr is required`;
    case value.length < 3:
      return `Accounr is too short`;
    case value.length > 100:
      return `Accounr is too long`;
    default:
      return "";
  }
};
