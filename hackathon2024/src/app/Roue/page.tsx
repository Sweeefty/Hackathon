'use client'

import { Layout } from "@/components/layout";
import { Box, Button, Center, Flex, Heading } from "@chakra-ui/react";
import React, { useState } from "react";
import { Wheel } from "react-custom-roulette";

interface Tondeuse {  /* interface pour typer les options de la roue */
  option: string;
}

/*  creation de ce tableau pour creer les possibilités de la roue
 */const martin: Tondeuse[] = [
  { option: "Place Warehouse" },
  { option: "Pain au chocolat" },
  { option: "Place WEI Gratuite" },
  { option: "goodies Ynov" },
];

 /* creation de la roue */
export default function Chaise() {

  const [mustSpin, setMustSpin] = useState<boolean>(false);  /* states pour la roue qui permettent de la faire tourner */

  const [prizeNumber, setPrizeNumber] = useState<number>(0); /* prizeNumber est la state qui permet de choisir l'option gagnante à partir de chiffre */

  const PouletBasquaise = () => {  /* fonction qui permet de faire tourner la roue et de choisir aléatoirement une option */
    const newPrizeNumber = Math.floor(Math.random() * martin.length);
    setPrizeNumber(newPrizeNumber);
    setMustSpin(true);
  };

  return (
    <>
    <Layout>
      <Flex 
      align="center"
      justify="center"
      h="100vh"
      bg="#1b202b">
        <Center h="100vh" flexDirection="column">


          <Flex justify="center" align="center">
            <Heading as="h1" size="xl" mb={6}>
              Roulette quotidienne
            </Heading>
              <Box w="80%">
                <Wheel /*  la roue qui a ete importer de react-custom-roulette */
                  mustStartSpinning={mustSpin} /*  reglages global de la roue */
                  prizeNumber={prizeNumber}
                  data={martin}
                  outerBorderColor={"#666666"}
                  outerBorderWidth={10}
                  innerBorderColor={"#666666"}
                  radiusLineColor={"#666666"}
                  radiusLineWidth={1}
                  fontSize={15}
                  textColors={["#ffffff"]}
                  backgroundColors={["#70DECF", "#70DECF"]}
                  onStopSpinning={() => { 
                    setMustSpin(false);
                    console.log(martin[prizeNumber]);
                  }}
                />
              </Box>
              <Button mt={4} p={10} bg={'#70DECF'} onClick={PouletBasquaise}> 
              Tourner la roue
            </Button>
          </Flex>
          {!mustSpin && <Box fontSize={30} mt={4}>{martin[prizeNumber].option}</Box>} {/* // affichage du resultat de la roue */}

        </Center>
      </Flex>
    </Layout>
    </>
  );
}
