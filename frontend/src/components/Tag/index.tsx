import React from 'react';
import cc from 'classcat';

import styles from './styles.module.scss';

interface Props {
  onClick?: () => void;
  children: React.ReactNode;
  active?: boolean;
}

export default function Tag({ children, onClick, active }: Props) {
  return (
    <button
      type="button"
      className={cc({
        [styles.tag]: true,
        [styles.hasFunc]: !!onClick,
        [styles.active]: active,
      })}
      onClick={onClick}
    >
      {children}
    </button>
  );
}
