import React from 'react';
import cc from 'classcat';
import useIsScrollingUp from '@/hooks/useIsScrollingUp';
import Github from '@/images/github.svg';
import Discord from '@/images/discord.svg';
import LinkedIn from '@/images/linkedin.svg';

import styles from './styles.module.scss';
import { useWindowScroll } from 'react-use';

export default function Header() {
  const isScrollingUp = useIsScrollingUp();
  const { y } = useWindowScroll();

  return (
    <header
      className={cc({
        [styles.header]: true,
        [styles.hideHeader]: !isScrollingUp && y !== 0,
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
