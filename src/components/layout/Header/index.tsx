import React from 'react';
import DiscordLogo from '@/images/icon/discord.svg';
import GithubLogo from '@/images/icon/github.svg';
import LinkedInLogo from '@/images/icon/linkedIn.svg';

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
          <DiscordLogo />
          <LinkedInLogo />
          <GithubLogo />
        </div>
      </div>
    </header>
  );
};

export default Header;
