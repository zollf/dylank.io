import React from 'react';
import HomeIcon from '@/images/home.svg';

import styles from './styles.module.scss';

const HelloWorld = () => {
  return (
    <div className={styles.helloWorld}>
      <HomeIcon />
      Hello World
    </div>
  );
};

export default HelloWorld;
