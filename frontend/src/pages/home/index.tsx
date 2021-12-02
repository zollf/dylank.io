import React from 'react';
import Header from '@/components/layout/Header';
import Hero from '@/components/Hero';
import MobileHeader from '@/components/layout/MobileHeader';
import useIsMobile from '@/hooks/useIsMobile';

import styles from './styles.module.scss';

export default function Home() {
  const [isMobile] = useIsMobile(768);
  return (
    <div className={styles.index}>
      {isMobile ? <MobileHeader /> : <Header />}
      <Hero />
    </div>
  );
}
