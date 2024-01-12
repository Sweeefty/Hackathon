'use client'

import { Layout } from "@/components/layout";
import { Box, Button, Center, Flex, Heading } from "@chakra-ui/react";
import React, { useState } from "react";
import { Wheel } from "react-custom-roulette";

interface Option {
  option: string;
}

const data: Option[] = [
  { option: "Place Warehouse" },
  { option: "Pain au chocolat" },
  { option: "Place WEI Gratuite" },
  { option: "goodies Ynov" },
];

export default function RouletteWheel() {
  const [mustSpin, setMustSpin] = useState<boolean>(false);
  const [prizeNumber, setPrizeNumber] = useState<number>(0);

  const handleSpinClick = () => {
    const newPrizeNumber = Math.floor(Math.random() * data.length);
    setPrizeNumber(newPrizeNumber);
    setMustSpin(true);
  };

  return (
    <>
    <Layout>
      <Center h="100vh" flexDirection="column">
        <Heading as="h1" size="xl" mb={6}>
          Roulette quotidienne
        </Heading>

        <Flex justify="center" align="center">
          <Box w="80%">
            <Wheel
              mustStartSpinning={mustSpin}
              prizeNumber={prizeNumber}
              data={data}
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
                console.log(data[prizeNumber]);
              }}
            />
          </Box>
        </Flex>

        <Button mt={4} onClick={handleSpinClick}>
          tourner la roue
        </Button>
        {!mustSpin && <Box mt={4}>{data[prizeNumber].option}</Box>}
      </Center>
    </Layout>
    </>
  );
}
