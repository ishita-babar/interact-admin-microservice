import getHandler from '@/handlers/get_handler';
import Toaster from '@/utils/toaster';
import React, { useEffect, useState } from 'react';
import { SERVER_ERROR } from '@/config/errors';
import Loader from '@/components/loader';
import { Project } from '@/types';
import InfiniteScroll from 'react-infinite-scroll-component';
import ProjectComponent from '@/components/project';

const Projects = () => {
  const [projects, setProjects] = useState<Project[]>([]);
  const [loading, setLoading] = useState(true);
  const [hasMore, setHasMore] = useState(true);
  const [page, setPage] = useState(1);

  const getProjects = () => {
    setLoading(true);
    const URL = `/flags/projects?page=${page}&limit=${10}`;
    getHandler(URL)
      .then(res => {
        if (res.statusCode === 200) {
          const addedProjects = [...projects, ...(res.data.projects || [])];
          if (addedProjects.length === projects.length) setHasMore(false);
          setProjects(addedProjects);
          setPage(prev => prev + 1);
          setLoading(false);
        } else {
          if (res.data.message) Toaster.error(res.data.message, 'error_toaster');
          else {
            Toaster.error(SERVER_ERROR, 'error_toaster');
          }
        }
      })
      .catch(err => {
        Toaster.error(SERVER_ERROR, 'error_toaster');
      });
  };

  useEffect(() => {
    getProjects();
  }, []);

  return (
    <div className="w-full h-full p-4">
      {loading ? (
        <div className="w-[45vw] mx-auto max-lg:w-[85%] max-md:w-full">
          <Loader />
        </div>
      ) : (
        <InfiniteScroll
          className="w-[45vw] mx-auto max-lg:w-[85%] max-md:w-screen flex flex-col gap-2 max-lg:px-4 pb-base_padding"
          dataLength={projects.length}
          next={getProjects}
          hasMore={hasMore}
          loader={<Loader />}
        >
          {projects.length === 0 ? (
            <div>No Flagged Projects.</div>
          ) : (
            projects.map(project => <ProjectComponent key={project.id} project={project} setProjects={setProjects} />)
          )}
        </InfiniteScroll>
      )}
    </div>
  );
};

export default Projects;
