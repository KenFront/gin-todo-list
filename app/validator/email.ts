import validator from "validator";

export const validateEmail = (value: string) => {
  switch (true) {
    case !value:
      return `Email is required`;
    case !validator.isEmail(value):
      return `Email is invalid`;
    default:
      return "";
  }
};
