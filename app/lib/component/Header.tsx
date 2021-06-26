import { ReactNode } from "react";
import {
  HStack,
  Button,
  useColorMode,
  Heading,
  Spacer,
} from "@chakra-ui/react";
import { MoonIcon, SunIcon } from "@chakra-ui/icons";
import styled from "@emotion/styled";

const CutomHStack = styled(HStack)`
  box-shadow: var(--chakra-shadows-sm);
  padding: 0 var(--chakra-space-6);
`;

export const Header = ({
  title,
  rightArea,
}: {
  title: string;
  rightArea?: ReactNode;
}) => {
  const { colorMode, toggleColorMode } = useColorMode();
  return (
    <CutomHStack w="100%" h="4.5rem">
      <Heading>{title}</Heading>
      <Spacer />
      <HStack spacing={4}>
        {rightArea}
        <Button onClick={toggleColorMode}>
          {colorMode === "light" ? <MoonIcon /> : <SunIcon />}
        </Button>
      </HStack>
    </CutomHStack>
  );
};
