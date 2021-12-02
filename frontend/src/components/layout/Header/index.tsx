import { Discord, Github, LinkedIn } from '@/images';

import React from 'react';
import styles from './styles.module.scss';

const Header = () => {
  return (
    <header className={styles.header}>
      <div className={styles.inner}>
        <div className={styles.logo}>d.io</div>
        <nav>
          <a href="#home">Home</a>
          <a href="#work">Work</a>
          <a href="#about">About</a>
          <a href="#contact">Contact</a>
        </nav>
        <div className={styles.socials}>
          <Discord />
          <LinkedIn />
          <Github />
        </div>
      </div>
    </header>
  );
};

export default Header;
