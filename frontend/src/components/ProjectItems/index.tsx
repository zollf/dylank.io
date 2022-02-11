import React, { useContext } from 'react';

import Project from '../Project';
import styles from './styles.module.scss';
import { ProjectsContext } from '../Projects';
export default function ProjectItems() {
  const { projects, loading } = useContext<ProjectsContext>(ProjectsContext);

  if (loading) {
    return <div className={styles.projects}>Loading</div>;
  }

  return (
    <div className={styles.projects}>
      {projects?.items?.map((project) => project && <Project project={project} key={project.id} />)}
    </div>
  );
}
