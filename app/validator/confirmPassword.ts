export const validateConfirmPassword = (ps: string) => (value: string) => {
  switch (true) {
    case !value:
      return `Confirm Password is required`;
    case ps !== value:
        return `Confirm Password is not match password`;
    default:
      return "";
  }
};
