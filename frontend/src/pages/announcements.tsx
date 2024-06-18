import getHandler from '@/handlers/get_handler';
import Toaster from '@/utils/toaster';
import React, { useEffect, useState } from 'react';
import { SERVER_ERROR } from '@/config/errors';
import Loader from '@/components/loader';
import { Announcement } from '@/types';
import InfiniteScroll from 'react-infinite-scroll-component';
import AnnouncementComponent from '@/components/announcement';
import BaseWrapper from '@/wrappers/base';
import Sidebar from '@/components/common/sidebar';
import MainWrapper from '@/wrappers/main';

const Announcements = () => {
  const [announcements, setAnnouncements] = useState<Announcement[]>([]);
  const [loading, setLoading] = useState(true);
  const [hasMore, setHasMore] = useState(true);
  const [page, setPage] = useState(1);

  const getAnnouncements = () => {
    setLoading(true);
    const URL = `/flags/announcements?page=${page}&limit=${10}`;
    getHandler(URL)
      .then(res => {
        if (res.statusCode === 200) {
          const addedAnnouncements = [...announcements, ...(res.data.announcements || [])];
          if (addedAnnouncements.length === announcements.length) setHasMore(false);
          setAnnouncements(addedAnnouncements);
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
    getAnnouncements();
  }, []);

  return (
    <BaseWrapper>
      <Sidebar index={8} />
      <MainWrapper>
    <div className="w-full h-full p-4">
      {loading ? (
        <div className="w-[45vw] mx-auto max-lg:w-[85%] max-md:w-full">
          <Loader />
        </div>
      ) : (
        <InfiniteScroll
          className="w-[45vw] mx-auto max-lg:w-[85%] max-md:w-screen flex flex-col gap-2 max-lg:px-4 pb-base_padding"
          dataLength={announcements.length}
          next={getAnnouncements}
          hasMore={hasMore}
          loader={<Loader />}
        >
          {announcements.length === 0 ? (
            <div>No Flagged Announcements.</div>
          ) : (
            announcements.map(announcement => <AnnouncementComponent key={announcement.id} announcement={announcement} setAnnouncements={setAnnouncements} />)
          )}
        </InfiniteScroll>
      )}
    </div>
    </MainWrapper>
    </BaseWrapper>
  );
};

export default Announcements;
