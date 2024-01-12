import { Layout } from "@/components/layout"
import React from "react";
import { Box, Flex, Heading, Text, SimpleGrid } from "@chakra-ui/react";

export default function Apropos(){
    return(
        <Layout>
            <Box fontWeight='bold' w="70%" ml="auto" mr="auto" bg="#70DECF" borderRadius={8} mt={8} textAlign="center" p={6} color="white">
                BDE de : Nantes
            </Box>
            <br />
            <Flex justifyContent="space-between">
                <Box fontWeight="bold" w="48%" bg="#70DECF" borderRadius={8} mr="auto" ml="auto" textAlign="center" p={8} color="white" fontSize="xl">
                Bienvenue chez Ynov Nantes - Découvre le BDE !<br /><br />

                    Chers Ynoviens,
                    Le Bureau des Étudiants d'Ynov Nantes, c'est l'endroit où ça se passe !<br /><br />

                    Qu'est-ce qu'on fait ?<br />
                    On s'assure que chaque étudiant passe un super moment à Ynov. Soirées, ateliers créatifs, et voyages, on organise tout pour rendre ton expérience inoubliable.<br /><br />

                    Nos valeurs :<br />
                    On croit en l'entraide, la créativité, et en faire équipe. Si tu aimes l'idée de contribuer à une communauté cool, le BDE est fait pour toi !<br /><br />

                    Comment nous rejoindre :<br />
                    Suis-nous sur les réseaux, viens à nos réunions, et apporte tes idées. Ensemble, on crée l'histoire Ynovienne.<br /><br />

                    Prêts pour une aventure mémorable ? Rejoins le BDE d'Ynov Nantes !<br /><br />

                    À bientôt,<br />
                    Le BDE d'Ynov Nantes
                </Box>

                <Box w="48%" fontWeight="bold" ml={8} bg="#70DECF" borderRadius={8} p={8} color="white" textAlign="center">
                    
                    <Text mb={2} color="white" fontSize="sm" as='u'>
                        Présentation des membres: <br /><br />
                        Président :<br />
                            Lucas Tremblay<br /><br />
                        Vice-Présidente - Événements :<br />
                            Camille Leroux<br /><br />
                        Vice-Président - Communication :<br />
                            Maxime Dubois<br /><br />
                        Secrétaire Général(e) :<br />
                            Emma Martin<br /><br />
                        Trésorier(ère) :<br />
                            Nicolas Lefèvre<br /><br />
                        Responsable des Événements Culturels :<br />
                            Léa Blanchard<br /><br />
                        Responsable des Soirées :<br />
                            Thomas Moreau<br /><br />
                        Responsable des Projets Artistiques :<br />
                            Juliette Rousseau<br /><br />
                        Responsable des Sports et Loisirs :<br />
                            Hugo Dubois<br /><br />
                        Responsable des Relations Étudiantes :<br />
                            Inès Gagnon
                    </Text>
                </Box>
            </Flex>
        </Layout>
    )
}