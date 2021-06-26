export const validatePassword = (value: string) => {
    switch (true) {
      case !value:
        return `Password is required`;
      case value.length < 3:
        return `Password is too short`;
      case value.length > 100:
        return `Password is too long`;
      default:
        return "";
    }
  };
  