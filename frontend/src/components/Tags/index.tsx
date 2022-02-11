import React, { useContext, useEffect, useRef, useState } from 'react';
import { useScroll } from 'react-use';
import Tag from '@/components/Tag';
import { ChevronLeft, ChevronRight } from '@/images';
import cc from 'classcat';
import styles from './styles.module.scss';
import { ProjectsContext } from '../Projects';

export default function Tags() {
  const [start, setStart] = useState(true);
  const [end, setEnd] = useState(false);
  const { projects, loading, setActiveFilter, clearFilters, activeFilters } =
    useContext<ProjectsContext>(ProjectsContext);
  const slider = useRef<HTMLDivElement>(null);
  const { x } = useScroll(slider);
  function next() {
    if (slider.current) {
      slider.current.scrollLeft += slider.current.clientWidth;
    }
  }

  function prev() {
    if (slider.current) {
      slider.current.scrollLeft -= slider.current.clientWidth;
    }
  }

  useEffect(() => {
    if (slider.current) {
      const maxScroll = slider.current.scrollWidth - slider.current.clientWidth;
      setEnd(Math.floor(x) === Math.floor(maxScroll));
      setStart(x === 0);
    }
  }, [x, slider]);

  if (loading) {
    return (
      <div className={styles.tagsWrapper}>
        <div className={styles.tags} ref={slider}>
          Loading
        </div>
      </div>
    );
  }

  return (
    <div className={styles.tagsWrapper}>
      <button className={cc({ [styles.prev]: true, [styles.hidden]: start })} onClick={prev}>
        <ChevronLeft />
      </button>
      <div className={styles.tags} ref={slider}>
        <Tag onClick={clearFilters} active={!activeFilters.length}>
          All ({projects?.total})
        </Tag>
        {projects?.tags?.map(
          (tag) =>
            tag && (
              <Tag onClick={() => setActiveFilter(tag.slug)} active={activeFilters.includes(tag.slug)} key={tag.id}>
                {tag.title} ({tag.count})
              </Tag>
            ),
        )}
      </div>
      <button className={cc({ [styles.next]: true, [styles.hidden]: end })} onClick={next}>
        <ChevronRight />
      </button>
    </div>
  );
}
