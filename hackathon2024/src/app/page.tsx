import { Button, ChakraProvider } from '@chakra-ui/react'
import Image from 'next/image'
import { Modal } from '@chakra-ui/react'

export default function Home() {
  return (
    <ChakraProvider>
        <Button className='w-full h-full bg-white'>
          police
        </Button>
      
    </ChakraProvider>
  )
}
