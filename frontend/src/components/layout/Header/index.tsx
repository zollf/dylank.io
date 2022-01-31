import React from 'react';
import cc from 'classcat';
import useIsScrollingUp from '@/hooks/useIsScrollingUp';
import { Discord, Github, LinkedIn } from '@/images';

import styles from './styles.module.scss';

export default function Header() {
  const isScrollingUp = useIsScrollingUp();
  const isTop = typeof window != 'undefined' ? window.pageYOffset === 0 : true;

  return (
    <header
      className={cc({
        [styles.header]: true,
        [styles.hideHeader]: !isScrollingUp && !isTop,
      })}
    >
      <div className={styles.inner}>
        <div className={styles.logo}>d.io</div>
        <nav>
          <a href="#home">Home</a>
          <a href="#work">Work</a>
          <a href="#about">About</a>
          <a href="#contact">Contact</a>
        </nav>
        <div className={styles.socials}>
          <a href="#">
            <Discord />
          </a>
          <a href="#">
            <LinkedIn />
          </a>
          <a href="#">
            <Github />
          </a>
        </div>
      </div>
    </header>
  );
}
