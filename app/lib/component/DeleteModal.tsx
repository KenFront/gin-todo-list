import { useRef, RefObject, FC } from "react";
import {
  AlertDialog,
  AlertDialogBody,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogContent,
  AlertDialogOverlay,
  Button,
} from "@chakra-ui/react";
import { FocusableElement } from "@chakra-ui/utils";

export const DeleteModal: FC<{
  isOpen: boolean;
  title: string;
  description: string;
  confirmText: string;
  cancelText: string;
  onConfirm: () => Promise<void>;
  onCancel: () => Promise<void>;
}> = ({
  isOpen,
  title,
  description,
  confirmText,
  cancelText,
  onConfirm,
  onCancel,
}) => {
  const cancelRef = useRef<RefObject<FocusableElement>>();

  return (
    <>
      <AlertDialog
        isOpen={isOpen}
        leastDestructiveRef={cancelRef.current}
        onClose={onCancel}
      >
        <AlertDialogOverlay>
          <AlertDialogContent>
            <AlertDialogHeader fontSize="lg" fontWeight="bold">
              {title}
            </AlertDialogHeader>

            <AlertDialogBody>{description}</AlertDialogBody>

            <AlertDialogFooter>
              <Button onClick={onCancel}>{cancelText}</Button>
              <Button colorScheme="red" onClick={onConfirm} ml={3}>
                {confirmText}
              </Button>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialogOverlay>
      </AlertDialog>
    </>
  );
};
