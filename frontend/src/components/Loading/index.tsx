import React from 'react';
import cc from 'classcat';
import Spinner, { Size } from '@atlaskit/spinner';

import styles from './styles.module.scss';

interface Props {
  size: Size;
  className?: string;
}

export default function Loading({ size, className }: Props) {
  return (
    <div
      className={cc({
        [styles.spinner]: true,
        [className || '']: !!className,
      })}
    >
      <Spinner size={size} />
    </div>
  );
}
