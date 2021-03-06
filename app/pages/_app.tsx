import type { AppProps } from "next/app";
import { ChakraProvider } from "@chakra-ui/react";
import 'react-virtualized/styles.css';
import '@/style/global.css';

const App = ({ Component, pageProps }: AppProps) => (
  <ChakraProvider>
    <Component {...pageProps} />
  </ChakraProvider>
);

export default App;
