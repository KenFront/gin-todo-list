export const isNotEmpty = (name: string) => (value: string) => {
  switch (true) {
    case !value:
      return `${name} is required`;
    default:
      return "";
  }
};
