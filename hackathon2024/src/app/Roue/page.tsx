
'use client'
import { useState } from 'react';
import { Flex, Button } from "@chakra-ui/react";
import { Winwheel } from 'react-winwheel';
import 'react-winwheel/dist/react-winwheel.css';

export default function RoulettePage(){
    const [wheelSpinning, setWheelSpinning] = useState(false);

  const wheelOptions = {
    numSegments: 8, 
    outerRadius: 200, 
    segments: [
      { fillStyle: "#70DECF", text: "Option 1" },
      { fillStyle: "#ffffff", text: "Option 2" },
      { fillStyle: "#70DECF", text: "Option 3" },
    ],
    animation: {
      type: 'spinToStop',
      duration: 5, 
      spins: 8, 
    },
  };

  const handleSpinClick = () => {
    if (!wheelSpinning) {
      setWheelSpinning(true);
      const wheel = new Winwheel(wheelOptions);

      wheel.startAnimation();
      
      setTimeout(() => {
        setWheelSpinning(false);
      }, wheelOptions.animation.duration * 1000);
    }
  };

  return (
    <Flex
      align="center"
      justify="center"
      h="100vh"
    >
      <div style={{ position: "relative" }}>
        <Winwheel {...wheelOptions} />
      </div>
      <Button onClick={handleSpinClick} mt={4}>
        Tourner la roue
      </Button>
    </Flex>
  );
};

