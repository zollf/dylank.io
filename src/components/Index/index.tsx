import React from 'react';
import useIsMobile from '@/hooks/useIsMobile';

import Header from '../layout/Header';
import Hero from '../Hero';
import MobileHeader from '../layout/MobileHeader';
import styles from './styles.module.scss';

const Index = () => {
  const [isMobile] = useIsMobile(768);
  return (
    <div className={styles.index}>
      {isMobile ? <MobileHeader /> : <Header />}
      <Hero />
    </div>
  );
};

export default Index;
