import { Button, useColorMode, Heading } from "@chakra-ui/react";
import { MoonIcon, SunIcon } from "@chakra-ui/icons";
import styled from "@emotion/styled";

const Container = styled.header`
  height: 4.5rem;
  width: 100%;
  box-shadow: var(--chakra-shadows-sm);
  padding: 0 var(--chakra-space-6);
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

export const Header = ({ title }: { title: string }) => {
  const { colorMode, toggleColorMode } = useColorMode();
  return (
    <Container>
      <Heading>{title}</Heading>
      <Button onClick={toggleColorMode}>
        {colorMode === "light" ? <MoonIcon /> : <SunIcon />}
      </Button>
    </Container>
  );
};
