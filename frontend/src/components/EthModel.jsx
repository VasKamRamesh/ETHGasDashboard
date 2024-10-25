// src/components/EthModel.jsx
import React, { useRef } from "react";
import { Canvas, useFrame, useLoader } from "@react-three/fiber";
import { GLTFLoader } from "three/examples/jsm/loaders/GLTFLoader"; // Import GLTFLoader
import { OrbitControls } from "@react-three/drei"; // For camera controls

const EthereumDiamond = () => {
  const meshRef = useRef(); // Reference to the loaded object

  // Load the GLTF/GLB model
  const { scene } = useLoader(GLTFLoader, "/ethereum_logo/scene.gltf"); // Adjust the path to your .gltf/.glb file

  // Rotate the model in the clockwise direction (along the Y-axis)
  useFrame(() => {
    if (meshRef.current) {
      meshRef.current.rotation.y -= 0.01; // Spin clockwise
    }
  });

  return (
    <>
      <primitive ref={meshRef} object={scene} scale={2} position={[0, -1, 0]} />
    </>
  );
};

export default function DiamondViewer() {
  return (
    <Canvas camera={{ position: [0, 2, 5], near: 0.1, far: 10000 }}>
      <ambientLight intensity={0.5} />
      <directionalLight position={[5, 5, 5]} intensity={1} />
      <OrbitControls
        autoRotate
        autoRotateSpeed={1.0}
        minDistance={5}
        maxDistance={20}
      />
      <EthereumDiamond />
    </Canvas>
  );
}
