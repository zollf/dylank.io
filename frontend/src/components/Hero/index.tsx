import React from 'react';
import Circle from '@/images/circle.svg';
import Space from '@/images/space.svg';
import useIsMobile from '@/hooks/useIsMobile';

import Button from '../Button';
import styles from './styles.module.scss';

const Hero = () => {
  const [isMobile] = useIsMobile(768);

  return (
    <div className={styles.hero}>
      <div className={styles.background}>
        <Space />
      </div>
      <div className={styles.content}>
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
            Contact
          </Button>
          <Button size="large" theme="white" href="#work" full={isMobile}>
            Explore
          </Button>
        </div>
      </div>
      {!isMobile && (
        <>
          <div className={styles.bubble1}>
            <Circle />
          </div>
          <div className={styles.bubble2}>
            <Circle />
          </div>
          <div className={styles.bubble3}>
            <Circle />
          </div>
          <div className={styles.bubble4}>
            <Circle />
          </div>
        </>
      )}
    </div>
  );
};

export default Hero;
