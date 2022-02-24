import React, { useContext, useEffect, useRef, useState } from 'react';
import { useScroll } from 'react-use';
import Tag from '@/components/Tag';
import { ChevronLeft, ChevronRight } from '@/images';
import cc from 'classcat';
import styles from './styles.module.scss';
import { ProjectsContext } from '../Projects';
import Loading from '@/components/Loading';
import { sortBy } from 'lodash';

export default function Tags() {
  const [start, setStart] = useState(true);
  const [end, setEnd] = useState(false);
  const { projects, setActiveFilter, clearFilters, activeFilters } = useContext<ProjectsContext>(ProjectsContext);
  const slider = useRef<HTMLDivElement>(null);
  const { x } = useScroll(slider);
  function next() {
    if (slider.current) {
      slider.current.scrollLeft += slider.current.clientWidth * 0.8;
    }
  }

  function prev() {
    if (slider.current) {
      slider.current.scrollLeft -= slider.current.clientWidth * 0.8;
    }
  }

  useEffect(() => {
    if (slider.current) {
      const maxScroll = slider.current.scrollWidth - slider.current.clientWidth;
      setEnd(Math.floor(x) === Math.floor(maxScroll));
      setStart(x === 0);
    }
  }, [x, slider, projects?.tags]);

  return (
    <div className={styles.tagsWrapper}>
      <button className={cc({ [styles.prev]: true, [styles.hidden]: start })} onClick={prev}>
        <ChevronLeft />
      </button>
      <div className={styles.tags} ref={slider}>
        <Tag onClick={clearFilters} active={!activeFilters.length}>
          All ({projects?.total || 0})
        </Tag>
        {!!projects?.tags?.length ? (
          sortBy(projects.tags, ['count'])
            .reverse()
            .map((tag: Tag) => (
              <Tag onClick={() => setActiveFilter(tag.slug)} active={activeFilters.includes(tag.slug)} key={tag.id}>
                {tag.title} ({tag.count})
              </Tag>
            ))
        ) : (
          <Loading size="medium" className={styles.loading} />
        )}
      </div>
      <button className={cc({ [styles.next]: true, [styles.hidden]: end })} onClick={next}>
        <ChevronRight />
      </button>
    </div>
  );
}
