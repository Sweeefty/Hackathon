import React from "react";
import { Layout } from "@/components/layout";
import { Box, Flex, Heading, Text, SimpleGrid } from "@chakra-ui/react";

export default function Accueil() {
  return (
    <Layout>
      <Flex direction="column" align="center" justify="flex-start" h="100%">
        <Heading as="h1" size="xl" mt={4} mb={6}>
          BDE de : Nantes
        </Heading>

        <Flex align="center" justify="center" w="100%">
          <Box w="25%" p={4} bg="#70DECF" borderRadius={8} mb={4}>
            <Heading as="h2" size="md" mb={4} color="white">
              Profil
            </Heading>
            <Text fontWeight="bold" mb={2} color="white">
              Nom: Martin
            </Text>
            <Text fontWeight="bold" mb={2} color="white">
              Mail: martin@matin.com
            </Text>
            <Text fontWeight="bold" mb={2} color="white">
              Campus: Nantes
            </Text>
          </Box>

          <Box w="50%" ml={8} bg="#70DECF" borderRadius={8} p={8}>
            <div style={{ height: "100%" }}>Calendrier</div>
          </Box>
        </Flex>

        <Box mt={8} w="100%" textAlign="center">
          <Heading as="h2" size="lg" mb={40}>
            Prochains événements à venir
          </Heading>
          <SimpleGrid columns={[1, 2, 3]} spacing={16} mx="auto" maxW="800px">
            <Box bg="#70DECF" p={4} borderRadius={8} color="white">
              <Text fontWeight="bold">WareHouse</Text>
              <Text>Date : 01/04/2024</Text>
            </Box>
            <Box bg="#70DECF" p={4} borderRadius={8} color="white">
              <Text fontWeight="bold">WareHouse</Text>
              <Text>Date : 01/04/2024</Text>
            </Box>
            <Box bg="#70DECF" p={4} borderRadius={8} color="white">
              <Text fontWeight="bold">WareHouse</Text>
              <Text>Date : 01/04/2024</Text>
            </Box>
          </SimpleGrid>
        </Box>
      </Flex>
    </Layout>
  );
}
