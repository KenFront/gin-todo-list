import { useToast, UseToastOptions } from "@chakra-ui/react";

export const useAppToast = () => {
  const toast = useToast();
  const toastSuccess = (props?: UseToastOptions) =>
    toast({
      position: "top",
      status: "success",
      duration: 3000,
      isClosable: true,
      ...props,
    });
  const toastError = (props?: UseToastOptions) =>
    toast({
      position: "top",
      status: "error",
      duration: 3000,
      isClosable: true,
      ...props,
    });

  return {
    toastSuccess,
    toastError
  };
};
