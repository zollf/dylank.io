import React from 'react';
import Button from '@/components/Button';
import Github from '@/images/icons/github';
import View from '@/images/icons/view';

import styles from './styles.module.scss';

interface Props {
  project: Project;
}

export default function Project({ project }: Props) {
  return (
    <div className={styles.project}>
      {/* {!!project.assets?.length ? (
        <img className={styles.image} src={project.assets[0]?.url} alt={project.assets[0]?.slug} />
      ) : (
      )} */}
      <div className={styles.stub}></div>
      <div className={styles.title}>{project.title}</div>
      <div className={styles.desc}>{project.shortDescription}</div>
      <div className={styles.buttons}>
        {project.previewLink && (
          <Button href={project.previewLink} size="small" theme="blue" icon>
            View
            <View />
          </Button>
        )}
        {project.gitLink && (
          <Button href={project.gitLink} size="small" theme="white" icon>
            Repo
            <Github />
          </Button>
        )}
      </div>
    </div>
  );
}
