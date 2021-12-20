import React, { Suspense } from 'react';
import useIsMobile from '@/hooks/useIsMobile';
import { Canvas } from '@react-three/fiber';

import Button from '../Button';
import HeroModel from '../HeroModel';
import styles from './styles.module.scss';

const Hero = () => {
  const [isMobile] = useIsMobile(768);

  return (
    <div className={styles.hero}>
      <div className={styles.left}>
        <Canvas className={styles.model}>
          <Suspense fallback={null}>
            <HeroModel />
          </Suspense>
        </Canvas>
        <div className={styles.background} />
      </div>
      <div className={styles.right}>
        <h2>software developer</h2>
        <h1>Hi, I'm Dylan</h1>
        {!isMobile && (
          <p>
            Expertise in everything web related software. From frontend to backend, infrastructure that's concrete and
            beautifully displayed. Studying at the University of Western Australia in both Computer Science and Data
            Science. Have experience in working with big tech stacks and cooperating in teams for huge projects.
          </p>
        )}
        <div className={styles.btns}>
          <Button size="large" theme="blue" href="#contact" full={isMobile}>
            Contact me
          </Button>
          <Button size="large" theme="white" href="#work" full={isMobile}>
            Check my work
          </Button>
        </div>
      </div>
    </div>
  );
};

export default Hero;
