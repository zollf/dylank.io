import React from 'react';
import useIsMobile from '@/hooks/useIsMobile';

import Projects from '../Projects';
import styles from './styles.module.scss';

export default function Work() {
  const [isMobile] = useIsMobile(768);

  return (
    <div className={styles.work}>
      <div className={styles.top}>
        <h1>Explore My Projects</h1>
        {!isMobile && (
          <p>
            Check my commercial and non commercial projects. If you have any questions feel free to ask me for more
            information.
          </p>
        )}
      </div>
      <Projects />
    </div>
  );
}
