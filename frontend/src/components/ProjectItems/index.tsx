import React, { useContext } from 'react';

import Project from '../Project';
import styles from './styles.module.scss';
import { ProjectsContext } from '../Projects';

export default function ProjectItems() {
  const { projects } = useContext<ProjectsContext>(ProjectsContext);

  return (
    <div className={styles.projects}>
      {projects?.items?.map((project) => project && <Project project={project} key={project.id} />)}
    </div>
  );
}
