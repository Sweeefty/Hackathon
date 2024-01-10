import { Button, ChakraProvider, Center, Box, Flex, Heading, Input, Stack } from '@chakra-ui/react'
import Image from 'next/image'
import { Modal } from '@chakra-ui/react'

export default function Home() {
  return (
    <ChakraProvider>
      <Flex
      align="center"
      justify="center"
      h="100vh"
      bg="white" 
    >
      <Box p={12} borderWidth={1} borderRadius={12} boxShadow="lg" bg="#70DECF">
        <Heading as="h2" size="lg" mb={4}>
          Connexion
        </Heading>
        <Stack spacing={16}>
          <Stack spacing={6}>
            <Input placeholder='Email' size='sm' bg='white' />
            <Input placeholder='Mot de Passe' size='sm' bg='white' />
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
 