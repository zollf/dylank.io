import React from 'react';
import Button from '@/components/Button';
import Github from '@/images/icons/github';
import Tag from '@/components/Tag';
import View from '@/images/icons/view';

import styles from './styles.module.scss';

interface Props {
  project: Project;
}

export default function Project({ project }: Props) {
  const extraCount = project.tags ? Math.max(project.tags.length - 2, 0) : 0;

  return (
    <div className={styles.project}>
      {!!project.assets?.length ? (
        <img className={styles.image} src={project.assets[0]?.url} alt={project.assets[0]?.slug} />
      ) : (
        <div className={styles.stub}></div>
      )}
      <div className={styles.title}>{project.title}</div>
      <div className={styles.tags}>
        {project?.tags?.slice(0, 2).map((tag) => (
          <Tag small key={`${project.id}-${tag?.id}`}>
            {tag?.title}
          </Tag>
        ))}
        {!!extraCount && <Tag small>{extraCount}+</Tag>}
        {!project?.tags?.length && <Tag small>None</Tag>}
      </div>
      <div className={styles.desc}>{project.description}</div>
      <div className={styles.buttons}>
        {project.url && (
          <Button href={project.url} size="small" theme="blue" icon>
            <View /> Preview
          </Button>
        )}
        {project.git && (
          <Button href={project.git} size="small" theme="blue" icon>
            <Github /> Repo
          </Button>
        )}
      </div>
    </div>
  );
}
