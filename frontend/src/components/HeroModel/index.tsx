import React, { useRef } from 'react';
import * as three from 'three';
import { Environment, OrbitControls, PerspectiveCamera, Plane, Shadow } from '@react-three/drei';
import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader';
import { useFrame, useLoader } from '@react-three/fiber';

// istanbul ignore next
const HeroModel = () => {
  const obj = useRef<three.Mesh>();
  const gltf = useLoader(GLTFLoader, 'scene/scene.gltf');

  useFrame(() => (obj.current!.rotation.y += 0.001));

  return (
    <PerspectiveCamera>
      <OrbitControls enableDamping regress />
      <Environment preset="sunset" />
      <primitive object={gltf.scene} scale={1.8} position={[0.1, 1.4, 0]} ref={obj} />
      <Shadow opacity={0.5} scale={[6, 6, 6]} rotation-x={-Math.PI / 2} position={[0, -2.49, 0]} />
      <Plane args={[1000, 1000]} rotation={[-Math.PI / 2, 0, 0]} position={[0, -2.5, 0]}>
        <meshBasicMaterial attach="material" color="#116acc" />
      </Plane>
    </PerspectiveCamera>
  );
};

export default HeroModel;
