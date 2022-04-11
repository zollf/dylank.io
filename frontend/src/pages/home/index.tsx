import React from 'react';
import Contact from '@/components/Contact';
import Footer from '@/components/layout/Footer';
import Header from '@/components/layout/Header';
import Hero from '@/components/Hero';
import MobileHeader from '@/components/layout/MobileHeader';
import Work from '@/components/Work';
import useIsMobile from '@/hooks/useIsMobile';

import styles from './styles.module.scss';

export default function Home() {
  const [isMobile] = useIsMobile(768);
  return (
    <div className={styles.index}>
      <div className={styles.inner}>
        {isMobile ? <MobileHeader /> : <Header />}
        <Hero />
        <Work />
        <Contact />
        <Footer />
      </div>
    </div>
  );
}
