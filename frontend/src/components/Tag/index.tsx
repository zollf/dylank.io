import React from 'react';
import cc from 'classcat';

import styles from './styles.module.scss';

interface Props {
  onClick?: () => void;
  children: React.ReactNode;
  small?: boolean;
  active?: boolean;
}

export default function Tag({ children, onClick, small, active }: Props) {
  return (
    <button
      type="button"
      className={cc({
        [styles.tag]: true,
        [styles.small]: small,
        [styles.hasFunc]: !!onClick,
        [styles.active]: active,
      })}
      onClick={onClick}
    >
      {children}
    </button>
  );
}
