import { ChakraProvider, Box, Flex, Heading, Button, Spacer } from '@chakra-ui/react';
import NextLink from 'next/link';

export default function Header() {
  return (
    <Flex align="center" justify="space-between" p={4} bg="#70DECF">
      <Box>
        <Heading as="h1" size="lg" color="white">
          BDE Ynov
        </Heading>
      </Box>
      <Flex align="center">
        <Button as={NextLink} href="/Accueil" variant="ghost" color="#666666" mr={2}>
          Accueil
        </Button>
        <Button as={NextLink} href="/Roue" variant="ghost" color="#666666" mr={2}>
          Roue
        </Button>
        <Button as={NextLink} href="/Apropos" variant="ghost" color="#666666" mr={2}>
          À propos
        </Button>
        <Spacer />
        <Button as={NextLink} href="/" variant="ghost" color="#666666">
          Déconnexion
        </Button>
      </Flex>
    </Flex>
  );
};

export function Layout({ children }) {
  return (
    <ChakraProvider>
      <Header />
      <Box p={4}>{children}</Box>
    </ChakraProvider>
  );
};
