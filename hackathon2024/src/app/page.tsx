import { Button, ChakraProvider, Center, Box, Flex, Heading, Input, Stack } from '@chakra-ui/react'
import Image from 'next/image'

export default function Home() {
  return (
    <ChakraProvider>
      <Flex
      align="center"
      justify="center"
      h="100vh"
      bg="white" 
    >

      <Box position="absolute" top="4" left="4">
          <Image src="/logo.png" alt="YNOV Logo" width={300} height={300} />
        </Box>
      <Box p={28} borderWidth={1} borderRadius={10} boxShadow="lg" bg="#70DECF">
        <Heading as="h2" size="lg" mb={16}>
          Connexion
        </Heading>
        <Stack spacing={16}>
          <Stack spacing={6}>
            <Input placeholder='Email' size='sm' bg='#666666' />
            <Input placeholder='Mot de Passe' size='sm' bg='#666666' />
          </Stack>
          <Button bg='#666666' size="md">
            Connexion
          </Button>
        </Stack>
      </Box>
    </Flex>
      
    </ChakraProvider>
  )
}
 