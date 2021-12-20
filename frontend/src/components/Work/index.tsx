import React, { useState } from 'react';
import Project from '@/components/Project';
import Tag from '@/components/Tag';
import useIsMobile from '@/hooks/useIsMobile';
import { useQuery } from '@apollo/client';
import { xor } from 'lodash';

import query from './query.graphql';
import styles from './styles.module.scss';

interface WorkQuery {
  tags: Array<Tag>;
  projects: Array<Project>;
}

export default function Work() {
  const [activeFilters, setActiveFilters] = useState<Array<string>>(['all']);
  const [isMobile] = useIsMobile(768);
  const { data, loading, error } = useQuery<WorkQuery>(query);

  if (!data || loading || error) {
    return null;
  }

  const setActiveFilter = (name: string) => {
    setActiveFilters(xor([name], activeFilters));
  };

  return (
    <div className={styles.work}>
      <div className={styles.top}>
        <h1>Work</h1>
        {!isMobile && (
          <p>
            Check my commercial and non commercial projects. If you have any questions feel free to ask me for more
            information.
          </p>
        )}
      </div>
      <div className={styles.tags}>
        <Tag onClick={() => setActiveFilter('all')} active={activeFilters.includes('all')}>
          All ({data.projects.length})
        </Tag>
        {data.tags.map((tag) => (
          <Tag onClick={() => setActiveFilter(tag.id)} active={activeFilters.includes(tag.id)}>
            {tag.title} ({tag.count})
          </Tag>
        ))}
      </div>
      <div className={styles.projects}>
        {data.projects.map((project) => (
          <Project project={project} />
        ))}
      </div>
    </div>
  );
}
