import React from 'react';
import DiscordLogo from '@/images/icon/discord.svg';
import GithubLogo from '@/images/icon/github.svg';
import LinkedInLogo from '@/images/icon/linkedIn.svg';

import styles from './styles.module.scss';

interface Props {
  active: boolean;
}

const MobileMenu = ({ active }: Props) => {
  return (
    <div className={styles.mobileMenu} data-active={active} data-testid="mobileMenu">
      <div className={styles.inner}>
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
    </div>
  );
};

export default MobileMenu;
